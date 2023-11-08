package common

import (
	"reflect"
	"testing"
)

// TestUniqueSlice tests the UniqueSlice function.
func TestUniqueSlice(t *testing.T) {
	tests := []struct {
		slice     []int
		wantSlice []int
	}{
		{[]int{1, 2, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
		{[]int{1, 1, 1, 1}, []int{1}},
	}

	for _, test := range tests {
		if got := UniqueSlice(test.slice); !reflect.DeepEqual(got, test.wantSlice) {
			t.Errorf("UniqueSlice(%v) = %v, want %v", test.slice, got, test.wantSlice)
		}
	}
}

// TestSliceContains tests the SliceContains function.
func TestSliceContains(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	if !SliceContains(slice, 3) {
		t.Errorf("SliceContains(%v, %v) = false, want true", slice, 3)
	}

	if SliceContains(slice, 6) {
		t.Errorf("SliceContains(%v, %v) = true, want false", slice, 6)
	}
}

// TestSliceIntersection tests the SliceIntersection function.
func TestSliceIntersection(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{3, 4, 5}
	want := []int{3}

	got := SliceIntersection(slice1, slice2)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SliceIntersection(%v, %v) = %v, want %v", slice1, slice2, got, want)
	}
}

// SliceFindFieldValue tests the SliceFindFiledValue function.
func TestSliceFindFieldValue(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}

	testSlice := []testStruct{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 25},
	}

	success, err := SliceFindFieldValue(testSlice, "Age", 25)
	if err != nil {
		t.Errorf("SliceFindFieldValue returned an unexpected error: %v", err)
	}

	wantSuccess := []interface{}{testSlice[0], testSlice[2]}
	if !reflect.DeepEqual(success, wantSuccess) {
		t.Errorf("SliceFindFieldValue(%v, \"Age\", 25) = %v, want %v", testSlice, success, wantSuccess)
	}

	empty, err := SliceFindFieldValue(testSlice, "Nonexistentfield", 25)
	if err != nil {
		t.Errorf("SliceFindFieldValue returned an unexpected error: %v", err)
	}

	wantEmpty := []interface{}{}
	if len(wantEmpty) != 0 && len(empty) != 0 {
		t.Errorf("SliceFindFieldValue(%v, \"Nonexistentfield\", 25) = %v, want %v", testSlice, empty, wantEmpty)
	}
}

func TestGetCurrentCPUInfo(t *testing.T) {
	_, err := GetCurrentCPUInfo()
	if err != nil {
		t.Errorf("GetCurrentCPUInfo() error = %v", err)
	}
}

func TestGetCurrentMemoryInfo(t *testing.T) {
	_, err := GetCurrentMemoryInfo()
	if err != nil {
		t.Errorf("GetCurrentMemoryInfo() error = %v", err)
	}
}

func TestGetCurrentDiskUsage(t *testing.T) {
	_, err := GetCurrentDiskUsage("/")
	if err != nil {
		t.Errorf("GetCurrentDiskUsage() error = %v", err)
	}
}

func TestGetDiskPartitions(t *testing.T) {
	_, err := GetDiskPartitions(true)
	if err != nil {
		t.Errorf("GetDiskPartitions() error = %v", err)
	}
}

func TestGetSystemLoadAverage(t *testing.T) {
	_, err := GetSystemLoadAverage()
	if err != nil {
		t.Errorf("GetSystemLoadAverage() error = %v", err)
	}
}

func TestGetHostInfo(t *testing.T) {
	_, err := GetHostInfo()
	if err != nil {
		t.Errorf("GetHostInfo() error = %v", err)
	}
}

func TestGetNetworkInterfaces(t *testing.T) {
	_, err := GetNetworkInterfaces()
	if err != nil {
		t.Errorf("GetNetworkInterfaces() error = %v", err)
	}
}

func TestGetNetworkIOCounters(t *testing.T) {
	_, err := GetNetworkIOCounters(true)
	if err != nil {
		t.Errorf("GetNetworkIOCounters() error = %v", err)
	}
}

func TestGetInternalIPv4(t *testing.T) {
	ip, err := GetInternalIPv4()
	if err != nil {
		t.Errorf("GetInternalIPv4() error = %v", err)
	}
	if ip[:7] != "192.168" {
		t.Errorf("GetInternalIPv4() got = %v, want prefix = %v", ip, "192.168")
	}
}
