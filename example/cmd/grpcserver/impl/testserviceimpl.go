package impl

/*
 * Copyright 2020-2021 Aldelo, LP
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	util "github.com/aldelo/common"
	testpb "github.com/aldelo/connector/example/proto/test"
	"context"
)

type TestServiceImpl struct{
	testpb.UnimplementedAnswerServiceServer
}

// implemented service greeting
func (*TestServiceImpl) Greeting(ctx context.Context, q *testpb.Question) (*testpb.Answer, error) {
	s := q.Question + " = " + "Reply OK"
	return &testpb.Answer{
		Answer: s,
	}, nil
}

type TestStreamServiceImpl struct {
	testpb.UnimplementedAnswerServerStreamServiceServer
}

func (*TestStreamServiceImpl) StreamGreeting(in *testpb.Question, stream testpb.AnswerServerStreamService_StreamGreetingServer) error {
	for i := 0; i < 50; i++ {
		if err := stream.Send(&testpb.Answer{
			Answer: util.CurrentDateTime(),
		}); err != nil {
			return err
		}
	}

	return nil
}
