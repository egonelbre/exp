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

VkDeviceCreateInfo * createInfo() {
	VkDeviceCreateInfo *info = malloc(sizeof(VkDeviceCreateInfo));
	VkDeviceQueueCreateInfo *queue = malloc(sizeof(VkDeviceQueueCreateInfo));
	info->pQueueCreateInfos = queue;
	return info;
}

void create1(VkDeviceCreateInfo *x0) {}
void create2(VkDeviceCreateInfo *x0, VkDeviceCreateInfo *x1) {}
void create3(VkDeviceCreateInfo *x0, VkDeviceCreateInfo *x1, VkDeviceCreateInfo *x2) {}
void create4(VkDeviceCreateInfo *x0, VkDeviceCreateInfo *x1, VkDeviceCreateInfo *x2, VkDeviceCreateInfo *x3) {}
void create8(
	VkDeviceCreateInfo *x0, VkDeviceCreateInfo *x1, VkDeviceCreateInfo *x2, VkDeviceCreateInfo *x3,
	VkDeviceCreateInfo *x4, VkDeviceCreateInfo *x5, VkDeviceCreateInfo *x6, VkDeviceCreateInfo *x7
) {}
*/
import "C"

//go:noinline
func Nop() {}

func CNop() { C.nop() }

func NewDeviceCreateInfo() *C.VkDeviceCreateInfo {
	return C.createInfo()
}

func Release(xs ...*C.VkDeviceCreateInfo)  {
	for _, x := range xs {
		C.free(unsafe.Pointer(x))
	}
}

//go:noinline
func Args1(x0 *C.VkDeviceCreateInfo) {	}

//go:noinline
func Args2(x0, x1 *C.VkDeviceCreateInfo) {}

//go:noinline
func Args3(x0, x1, x2 *C.VkDeviceCreateInfo) {}

//go:noinline
func Args4(x0, x1, x2, x3 *C.VkDeviceCreateInfo) {}

//go:noinline
func Args8(x0, x1, x2, x3, x4, x5, x6, x7 *C.VkDeviceCreateInfo) {}

func CArgs1(x0 *C.VkDeviceCreateInfo) {	
	C.create1(x0)
}

func CArgs2(x0, x1 *C.VkDeviceCreateInfo) {
	C.create2(x0, x1)
}

func CArgs3(x0, x1, x2 *C.VkDeviceCreateInfo) {
	C.create3(x0, x1, x2)
}

func CArgs4(x0, x1, x2, x3 *C.VkDeviceCreateInfo) {
	C.create4(x0, x1, x2, x3)
}

func CArgs8(x0, x1, x2, x3, x4, x5, x6, x7 *C.VkDeviceCreateInfo) {
	C.create8(x0, x1, x2, x3, x4, x5, x6, x7)
}

