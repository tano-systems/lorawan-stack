// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redis_test

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/events/redis"
	ttnredis "go.thethings.network/lorawan-stack/v3/pkg/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

// redisConfig returns a new redis config for testing
func redisConfig() ttnredis.Config {
	var err error
	config := ttnredis.Config{
		Address:       "localhost:6379",
		Database:      1,
		RootNamespace: []string{"test"},
	}
	if address := os.Getenv("REDIS_ADDRESS"); address != "" {
		config.Address = address
	}
	if db := os.Getenv("REDIS_DB"); db != "" {
		config.Database, err = strconv.Atoi(db)
		if err != nil {
			panic(err)
		}
	}
	if prefix := os.Getenv("REDIS_PREFIX"); prefix != "" {
		config.RootNamespace = []string{prefix}
	}
	return config
}

func Example() {
	// The task starter is used for automatic re-subscription on failure.
	taskStarter := component.StartTaskFunc(component.DefaultStartTask)
	// This sends all events received from Redis to the default pubsub.
	redisPubSub := redis.WrapPubSub(context.TODO(), events.DefaultPubSub(), taskStarter, ttnredis.Config{
		// Config here...
	})
	// Replace the default pubsub so that we will now publish to Redis.
	events.SetDefaultPubSub(redisPubSub)
}

var timeout = (1 << 10) * test.Delay

func TestRedisPubSub(t *testing.T) {
	test.RunTest(t, test.TestConfig{
		Timeout: 4 * timeout,
		Func: func(ctx context.Context, a *assertions.Assertion) {
			events.IncludeCaller = true

			eventCh := make(chan events.Event)
			handler := events.HandlerFunc(func(e events.Event) {
				t.Logf("Received event %v", e)
				a.So(e.Time().IsZero(), should.BeFalse)
				a.So(e.Context(), should.NotBeNil)
				eventCh <- e
			})

			ctx = events.ContextWithCorrelationID(ctx, t.Name())

			taskStarter := component.StartTaskFunc(component.DefaultStartTask)
			pubsub := redis.NewPubSub(ctx, taskStarter, redisConfig())
			defer pubsub.Close(ctx)

			pubsub.Subscribe("redis.**", handler)
			time.Sleep(timeout)

			appID := &ttnpb.ApplicationIdentifiers{ApplicationID: "test-app"}

			test.RunSubtestFromContext(ctx, test.SubtestConfig{
				Name:    "publish",
				Timeout: timeout,
				Func: func(ctx context.Context, t *testing.T, a *assertions.Assertion) {
					pubsub.Publish(events.New(ctx, "redis.test.evt0", "redis test event 0", events.WithIdentifiers(appID)))
					select {
					case e := <-eventCh:
						a.So(e.Name(), should.Equal, "redis.test.evt0")
						if a.So(e.Identifiers(), should.NotBeNil) && a.So(e.Identifiers(), should.HaveLength, 1) {
							a.So(e.Identifiers()[0].GetApplicationIDs(), should.Resemble, appID)
						}
					case <-ctx.Done():
						t.Error("Did not receive expected event")
						t.FailNow()
					}
				},
			})
		},
	})
}
