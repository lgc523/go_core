package channel

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type ReusableObj struct {
}
type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow!")
	}
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
		t.Error(err)
	}
	for i := 0; i < 11; i++ {
		if obj, err := pool.GetObj(time.Millisecond * 50); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("addr %T\n", obj)
			if err := pool.ReleaseObj(obj); err != nil {
				t.Error(err)
			}
		}
	}
}
