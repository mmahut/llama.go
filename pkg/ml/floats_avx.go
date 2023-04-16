//go:build !noasm && amd64

// AUTO-GENERATED BY GOAT -- DO NOT EDIT

package ml

import "unsafe"

//go:noescape
func _mm256_mul_const_add_to(a, b, c, n unsafe.Pointer)

//go:noescape
func _mm256_mul_const_to(a, b, c, n unsafe.Pointer)

//go:noescape
func _mm256_mul_const(a, b, n unsafe.Pointer)

//go:noescape
func _mm256_mul_to(a, b, c, n unsafe.Pointer)

//go:noescape
func _mm256_dot(a, b, n, ret unsafe.Pointer)
