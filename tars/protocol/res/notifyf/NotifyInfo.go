//
// This file war generated by FastTars2go 1.0
// Generated from NotifyF.tars
// Tencent.

package notifyf

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

type NotifyInfo struct {
	Nextpage    int32        `json:"nextpage"`
	NotifyItems []NotifyItem `json:"notifyItems"`
}

func (st *NotifyInfo) resetDefault() {
}

func (st *NotifyInfo) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_int32(&st.Nextpage, 1, true)
	if err != nil {
		return err
	}

	err, have, ty = _is.SkipToNoCheck(2, true)
	if err != nil {
		return err
	}

	if ty == codec.LIST {
		err = _is.Read_int32(&length, 0, true)
		if err != nil {
			return err
		}
		st.NotifyItems = make([]NotifyItem, length, length)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {

			err = st.NotifyItems[i0].ReadBlock(_is, 0, false)
			if err != nil {
				return err
			}
		}
	} else if ty == codec.SIMPLE_LIST {
		err = fmt.Errorf("type not support SIMPLE_LIST.")
		if err != nil {
			return err
		}
	} else {
		err = fmt.Errorf("require vector, but not.")
		if err != nil {
			return err
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

func (st *NotifyInfo) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require NotifyInfo, but not exist. tag %d", tag)
		} else {
			return nil
		}
	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

func (st *NotifyInfo) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(st.Nextpage, 1)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.LIST, 2)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(st.NotifyItems)), 0)
	if err != nil {
		return err
	}
	for _, v := range st.NotifyItems {

		err = v.WriteBlock(_os, 0)
		if err != nil {
			return err
		}
	}

	return nil
}

func (st *NotifyInfo) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}