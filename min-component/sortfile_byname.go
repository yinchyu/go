package main

import (
	"os"
	"path/filepath"
	"sort"
)




type (
	sortDirType uint8
	dirInfoSort struct {
		fi       []os.FileInfo
		sortType sortDirType
	}
)

func (d *dirInfoSort) Len() int {
	return len(d.fi)
}

func (d *dirInfoSort) Default(x, y int) bool {
	lx, ly := len(d.fi[x].Name()), len(d.fi[y].Name())
	if lx == ly {
		return d.fi[x].Name() < d.fi[y].Name()
	}
	return lx < ly
}

const (
	sortDirTypeByNameAsc sortDirType = iota
	sortDirTypeByNameDesc
	sortDirTypeByTimeAsc
	sortDirTypeByTimeDesc
	sortDirTypeBySizeAsc
	sortDirTypeBySizeDesc
	sortDirTypeByExtAsc
	sortDirTypeByExtDesc
)

func (d *dirInfoSort) Less(x, y int) bool {
	if d.fi[x].IsDir() != d.fi[y].IsDir() {
		return d.fi[x].IsDir() // 文件夹永远排在文件前面
	}
	switch d.sortType {
	default:
		fallthrough // 默认使用文件名升序
	case sortDirTypeByNameAsc:
		return d.Default(x, y)
	case sortDirTypeByNameDesc:
		// 两个string 可以直接进行比较
		lx, ly := len(d.fi[x].Name()), len(d.fi[y].Name())
		if lx == ly {
			return d.fi[x].Name() > d.fi[y].Name()
		}
		return lx > ly
	case sortDirTypeByTimeAsc:
		tx, ty := d.fi[x].ModTime(), d.fi[y].ModTime()
		if tx.Unix() == ty.Unix() {
			return d.Default(x, y)
		}
		return tx.Before(ty)
	case sortDirTypeByTimeDesc:
		tx, ty := d.fi[x].ModTime(), d.fi[y].ModTime()
		if tx.Unix() == ty.Unix() {
			return d.Default(x, y)
		}
		return tx.After(ty)
	case sortDirTypeBySizeAsc:
		sx, sy := d.fi[x].Size(), d.fi[y].Size()
		if sx == sy {
			return d.Default(x, y)
		}
		return sx < sy
	case sortDirTypeBySizeDesc:
		sx, sy := d.fi[x].Size(), d.fi[y].Size()
		if sx == sy {
			return d.Default(x, y)
		}
		return sx > sy
	case sortDirTypeByExtAsc:
		if !d.fi[x].IsDir() && !d.fi[y].IsDir() {
			// 按照文件的类型进行排序
			return filepath.Ext(d.fi[x].Name()) < filepath.Ext(d.fi[y].Name())
		}
		return d.Default(x, y)
	case sortDirTypeByExtDesc:
		if !d.fi[x].IsDir() && !d.fi[y].IsDir() {
			return filepath.Ext(d.fi[x].Name()) > filepath.Ext(d.fi[y].Name())
		}
		return d.Default(x, y)
	}
}

func (d *dirInfoSort) Swap(x, y int) {
	d.fi[x], d.fi[y] = d.fi[y], d.fi[x]
}

func sortDir(dir string, inputType int) ([]os.FileInfo, error) {
	sortType := sortDirType(inputType)
	if sortType < sortDirTypeByNameAsc || sortType > sortDirTypeByExtDesc {
		sortType = sortDirTypeByNameAsc
		//*inputType = int(sortType)
	}
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	_ = f.Close()
	if err != nil {
		return nil, err
	}
	sort.Sort(&dirInfoSort{fi: list, sortType: sortType})
	return list, nil
}
