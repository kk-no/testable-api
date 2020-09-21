package store_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/kk-no/testable-api/database"
	"github.com/kk-no/testable-api/models"
	"github.com/kk-no/testable-api/store"
	"github.com/kk-no/testable-api/test/testutils"
)

func TestUser_FindByID(t *testing.T) {
	testutils.SetupOptionalFixtures([]string{"Users.sql"})
	defer testutils.TruncateTables()

	u := store.NewUser(database.Conn)

	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				id:  "1",
			},
			want: &models.User{
				ID:   "1",
				Name: "tester1",
			},
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				id:  "2",
			},
			want: &models.User{
				ID:   "2",
				Name: "tester2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := u.FindByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
