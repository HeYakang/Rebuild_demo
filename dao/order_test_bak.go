package dao

import (
	"Rebuild_demo/dao/db"
	"Rebuild_demo/model"
	"gorm.io/gorm"
	"reflect"
	"testing"
)





func Test_orderDAO_Create(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		req *model.DemoOrder
	}
	Db := db.NewOrm()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name:"all right",fields:fields{db:Db},args:args{req: &model.DemoOrder{OrderNo: "1111",UserName: "heyakang",Amount: 1.233}},wantErr: false},
		{name:"loss name",fields:fields{db:Db},args:args{req: &model.DemoOrder{OrderNo: "2222",Amount: 11.233}},wantErr: false},
		{name:"loss Amount",fields:fields{db:Db},args:args{req: &model.DemoOrder{OrderNo: "3333",UserName: "heyakang12"}},wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &orderDAO{
				db: tt.fields.db,
			}
			if err := o.Create(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			} else{
			   result,_ := o.QueryByNo(tt.args.req.OrderNo)
			   t.Log(result)
			}
		})
	}
}

func Test_orderDAO_DeleteByNo(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		orderNo string
	}
	Db := db.NewOrm()
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantOrder *model.DemoOrder
		wantErr   bool
	}{
		// TODO: Add test cases.
		{name:"delete 1111",fields:fields{db:Db},args:args{orderNo: "1111"},wantErr: false},
		{name:"delete 2222",fields:fields{db:Db},args:args{orderNo: "2222"},wantErr: false},
		{name:"delete 3333",fields:fields{db:Db},args:args{orderNo: "3333"},wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &orderDAO{
				db: tt.fields.db,
			}
			 err := o.DeleteByNo(tt.args.orderNo)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}else{
				result,err := o.QueryByNo(tt.args.orderNo)
				if err != nil {
					t.Errorf("DeleteByID() error = %v, wantErr %v", err, tt.wantErr)
				}
				t.Log(result)
			}

			//if !reflect.DeepEqual(gotOrder, tt.wantOrder) {
			//	t.Errorf("DeleteByID() gotOrder = %v, want %v", gotOrder, tt.wantOrder)
			//}
		})
	}
}

func Test_orderDAO_QueryByNo(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		orderNo string
	}
	Db := db.NewOrm()
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantOrder *model.DemoOrder
		wantErr   bool
	}{
		// TODO: Add test cases.
		{name:"search 1111",fields:fields{db:Db},args:args{orderNo: "1111"},wantErr: false},
		{name:"search 2222",fields:fields{db:Db},args:args{orderNo: "2222"},wantErr: false},
		{name:"search 3333",fields:fields{db:Db},args:args{orderNo: "3333"},wantErr: false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &orderDAO{
				db: tt.fields.db,
			}
			gotOrder, err := o.QueryByNo(tt.args.orderNo)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryByNo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//这边可以写想要的结果作为对比 但是目前没写
			if !reflect.DeepEqual(gotOrder, tt.wantOrder) {
				t.Errorf("QueryByNo() gotOrder = %v, want %v", gotOrder, tt.wantOrder)
			}
		})
	}
}

func Test_orderDAO_QueryListByName(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		userName string
		orderBy  string
	}
	Db := db.NewOrm()
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOrders []model.DemoOrder
		wantErr    bool
	}{
		// TODO: Add test cases.
		{name:"search heyakang",fields:fields{db:Db},args:args{userName: "heyakang"},wantErr: false},
		{name:"search heyakang1",fields:fields{db:Db},args:args{userName: "heyakang1"},wantErr: false},
		{name:"search null",fields:fields{db:Db},args:args{userName: ""},wantErr: false},
		{name:"search heyakang1233 no result",fields:fields{db:Db},args:args{userName: "heyakang1233"},wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &orderDAO{
				db: tt.fields.db,
			}
			gotOrders, err := o.QueryListByName(tt.args.userName, tt.args.orderBy)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryListByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrders, tt.wantOrders) {
				t.Errorf("QueryListByName() gotOrders = %v, want %v", gotOrders, tt.wantOrders)
			}
		})
	}
}

func Test_orderDAO_QueryTable(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	Db := db.NewOrm()
	tests := []struct {
		name       string
		fields     fields
		wantOrders []model.DemoOrder
		wantErr    bool
	}{
		// TODO: Add test cases.
		//无参数测一个
		{name:"search 1111",fields:fields{db:Db},wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &orderDAO{
				db: tt.fields.db,
			}
			gotOrders, err := o.QueryTable()
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryTable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrders, tt.wantOrders) {
				t.Errorf("QueryTable() gotOrders = %v, want %v", gotOrders, tt.wantOrders)
			}
		})
	}
}

func Test_orderDAO_UpdateByNo(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		orderNo string
		m       map[string]interface{}
	}
	Db := db.NewOrm()
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name:"updata 1111 to 111111.1",fields:fields{db:Db},args:args{orderNo: "1111",m: map[string]interface{}{"amount":111111.1}},wantErr: false},
		{name:"updata 2222 to 222222.2",fields:fields{db:Db},args:args{orderNo: "2222",m: map[string]interface{}{"amount":222222.2}},wantErr: false},
		{name:"updata 3333 to 333333.3",fields:fields{db:Db},args:args{orderNo: "3333",m: map[string]interface{}{"amount":333333.3}},wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &orderDAO{
				db: tt.fields.db,
			}
			if err := o.UpdateByNo(tt.args.orderNo, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("UpdateByNo() error = %v, wantErr %v", err, tt.wantErr)
			}else{
				result,_ := o.QueryByNo(tt.args.orderNo)
				t.Log(result)
			}
		})
	}
}

