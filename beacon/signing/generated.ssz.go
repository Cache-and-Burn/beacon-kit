// Code generated by fastssz. DO NOT EDIT.
// Hash: a92b4c6fef5764f8be94d5ee97b9d560c7c28a7048701cdafcf70cf703490e0a
package signing

import (
	ssz "github.com/prysmaticlabs/fastssz"
)

// MarshalSSZ ssz marshals the Data object
func (d *Data) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(d)
}

// MarshalSSZTo ssz marshals the Data object to a target array
func (d *Data) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'ObjectRoot'
	if size := len(d.ObjectRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ObjectRoot", size, 32)
		return
	}
	dst = append(dst, d.ObjectRoot...)

	// Field (1) 'Domain'
	dst = append(dst, d.Domain[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the Data object
func (d *Data) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 64 {
		return ssz.ErrSize
	}

	// Field (0) 'ObjectRoot'
	if cap(d.ObjectRoot) == 0 {
		d.ObjectRoot = make([]byte, 0, len(buf[0:32]))
	}
	d.ObjectRoot = append(d.ObjectRoot, buf[0:32]...)

	// Field (1) 'Domain'
	copy(d.Domain[:], buf[32:64])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Data object
func (d *Data) SizeSSZ() (size int) {
	size = 64
	return
}

// HashTreeRoot ssz hashes the Data object
func (d *Data) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(d)
}

// HashTreeRootWith ssz hashes the Data object with a hasher
func (d *Data) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'ObjectRoot'
	if size := len(d.ObjectRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("--.ObjectRoot", size, 32)
		return
	}
	hh.PutBytes(d.ObjectRoot)

	// Field (1) 'Domain'
	hh.PutBytes(d.Domain[:])

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}
