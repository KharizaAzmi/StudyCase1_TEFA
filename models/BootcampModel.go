package models

import (
	"encoding/binary"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BootcampModel struct {
	ID            primitive.ObjectID `bson:"id_bc"`
	JudulBootcamp string  `bson:"bootcamp"`
	Deskripsi 	  string  `bson:"deskripsi"`
	Harga 		  float64 `bson:"harga"`
	Durasi 		  string     `bson:"durasi"`
	Jadwal 		  string  `bson:"jadwal"`
} 

type ContentBootcamp struct {
	ID            primitive.ObjectID `bson:"id_video"`
	JudulBootcamp *BootcampModel  `bson:"bootcamp"`
	Materi 		  string  `bson:"materi_bc"`
	Video		  binary.AppendByteOrder `bson:"video_bc"`
}

type UserBootcamp struct {
	ID			  primitive.ObjectID `bson:"id_video"`
	User    *UserModel `bson:"akun_user"`
	BootcampUser  *ContentBootcamp `bson:"BootcampUser"`
}