package tfidf_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"golang-dev-challenge-mid/pkg/tfidf"
	"testing"
)

func TestRun(t *testing.T) {

	tt := []struct {
		name    string
		address string
		expect  float64
		error   error
	}{
		{
			name:    "test tfidf success",
			address: "../../data",
			expect:  -0. - 0.2879076464761811,
			error:   nil,
		},
		{
			name:    "test tfidf failed",
			address: "../data",
			expect:  -0. - 0.2879076464761811,
			error:   errors.New("error file not found: open ../data: The system cannot find the file specified."),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			result, err := tfidf.Run(tc.address)

			if tc.error != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.error.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expect, result)
			}
		})
	}

}
