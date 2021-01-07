package db

import "github.com/tebrizetayi/cleanarchitecture/domain/model"

var once sync.Once

type AuthorsInmemory struct {
	Author map[int]model.Author
	sequence int
}

type ArticlesInmemory struct {
	Article map[int]model.Author
	sequence int
}

type Inmemorydb struct {
	Authors  AuthorsInmemory
	Articles ArticlesInmemory
}

var instance *Inmemorydb
func GetInmemorydb() *Inmemorydb {
	once.Do(func(){
		instance:=&Inmemorydb{
			Auhors:AuthorsInmemory{
				Authors: make(map[int]model.Author),
			},
			Articles:ArticleInmemory{
				Articles:make(map[int]model.Article),
			}
		}
	})
	return instance
}

