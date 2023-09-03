package call

import "unsafe"

/*
#include <stdint.h>
#include <stdlib.h>

// Vulkan like data structures

typedef uint32_t VkFlags;
typedef VkFlags  VkDeviceQueueCreateFlags;
typedef uint32_t VkStructureType;

typedef struct VkDeviceQueueCreateInfo {
    VkStructureType             sType;
    const void*                 pNext;
    VkDeviceQueueCreateFlags    flags;
    uint32_t                    queueFamilyIndex;
    uint32_t                    queueCount;
    const float*                pQueuePriorities;
} VkDeviceQueueCreateInfo;

typedef struct VkPhysicalDeviceFeatures {
    uint32_t bools[56];
} VkPhysicalDeviceFeatures;

typedef struct VkDeviceCreateInfo {
    VkStructureType                    sType;
    const void*                        pNext;
    VkFlags                            flags;
    uint32_t                           queueCreateInfoCount;
    const VkDeviceQueueCreateInfo*     pQueueCreateInfos;
    uint32_t                           enabledLayerCount;
    const char* const*                 ppEnabledLayerNames;
    uint32_t                           enabledExtensionCount;
    const char* const*                 ppEnabledExtensionNames;
    const VkPhysicalDeviceFeatures*    pEnabledFeatures;
} VkDeviceCreateInfo;

void nop() { }

int32_t add2(int32_t x, int32_t y) { return x + y; }

VkDeviceCreateInfo * createInfo() {
	VkDeviceCreateInfo *info = malloc(sizeof(VkDeviceCreateInfo));
	VkDeviceQueueCreateInfo *queue = malloc(sizeof(VkDeviceQueueCreateInfo));
	info->pQueueCreateInfos = queue;
	return info;
}

void create1(VkDeviceCreateInfo *x0) {}
void create2(VkDeviceCreateInfo *x0, VkDeviceCreateInfo *x1) {}
void create4(VkDeviceCreateInfo *x0, VkDeviceCreateInfo *x1, VkDeviceCreateInfo *x2, VkDeviceCreateInfo *x3) {}
*/
import "C"

//go:noinline
func Nop() {}

func CNop() { C.nop() }

func NewDeviceCreateInfo() *C.VkDeviceCreateInfo {
	return C.createInfo()
}

func Release(xs ...*C.VkDeviceCreateInfo) {
	for _, x := range xs {
		C.free(unsafe.Pointer(x))
	}
}

//go:noinline
func Add2(x, y int32) int32 { return x + y }

func CAdd2(x, y int32) int32 { return int32(C.add2(C.int32_t(x), C.int32_t(y))) }

//go:noinline
func Args1(x0 *C.VkDeviceCreateInfo) {}

//go:noinline
func Args2(x0, x1 *C.VkDeviceCreateInfo) {}

//go:noinline
func Args4(x0, x1, x2, x3 *C.VkDeviceCreateInfo) {}

func CArgs1(x0 *C.VkDeviceCreateInfo) {
	C.create1(x0)
}

func CArgs2(x0, x1 *C.VkDeviceCreateInfo) {
	C.create2(x0, x1)
}

func CArgs4(x0, x1, x2, x3 *C.VkDeviceCreateInfo) {
	C.create4(x0, x1, x2, x3)
}
