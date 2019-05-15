package sdl

/*
#include "sdl_wrapper.h"

#if (SDL_VERSION_ATLEAST(2,0,6))
#if defined(_WIN32)
	#include <SDL2/SDL_vulkan.h>
#else
	#include <SDL_vulkan.h>
#endif
#else //!SDL_VERSION_ATLEAST(2,0,6)

#define VK_DEFINE_HANDLE(object) typedef struct object##_T* object;
#if defined(__LP64__) || defined(_WIN64) || defined(__x86_64__) || defined(_M_X64) || defined(__ia64) || defined (_M_IA64) || defined(__aarch64__) || defined(__powerpc64__)
#define VK_DEFINE_NON_DISPATCHABLE_HANDLE(object) typedef struct object##_T *object;
#else
#define VK_DEFINE_NON_DISPATCHABLE_HANDLE(object) typedef uint64_t object;
#endif

VK_DEFINE_HANDLE(VkInstance)
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkSurfaceKHR)
typedef VkInstance SDL_vulkanInstance;
typedef VkSurfaceKHR SDL_vulkanSurface;


#if defined(WARN_OUTDATED)
#pragma message("SDL_Vulkan_LoadLibrary is not supported before SDL 2.0.6")
#endif

static int SDL_Vulkan_LoadLibrary(const char *path)
{
	return 0;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_Vulkan_GetVkGetInstanceProcAddr is not supported before SDL 2.0.6")
#endif

static void* SDL_Vulkan_GetVkGetInstanceProcAddr(void)
{
	return NULL;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_Vulkan_UnloadLibrary is not supported before SDL 2.0.6")
#endif

static void SDL_Vulkan_UnloadLibrary(void) {}


#if defined(WARN_OUTDATED)
#pragma message("SDL_Vulkan_GetInstanceExtensions is not supported before SDL 2.0.6")
#endif

static int SDL_Vulkan_GetInstanceExtensions(SDL_Window *window, unsigned int *pCount, const char **pNames)
{
	return 0;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_Vulkan_CreateSurface is not supported before SDL 2.0.6")
#endif

static int SDL_Vulkan_CreateSurface(SDL_Window *window, VkInstance instance, VkSurfaceKHR *surface)
{
	return 0;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_Vulkan_GetDrawableSize is not supported before SDL 2.0.6")
#endif

static void SDL_Vulkan_GetDrawableSize(SDL_Window *window, int *w, int *h) {
	*w = 0;
	*h = 0;
}

#endif
*/
import "C"
import (
	"errors"
	"reflect"
	"unsafe"
)

// VulkanLoadLibrary dynamically loads a Vulkan loader library.
// (https://wiki.libsdl.org/SDL_Vulkan_LoadLibrary)
func VulkanLoadLibrary(path string) error {
	var ret C.int
	if path == "" {
		ret = C.SDL_Vulkan_LoadLibrary(nil)
	} else {
		cpath := C.CString(path)
		defer C.free(unsafe.Pointer(cpath))
		ret = C.SDL_Vulkan_LoadLibrary(cpath)
	}
	if int(ret) == -1 {
		return GetError()
	}
	return nil
}

// VulkanGetVkGetInstanceProcAddr gets the address of the vkGetInstanceProcAddr function.
// (https://wiki.libsdl.org/SDL_Vulkan_GetVkInstanceProcAddr)
func VulkanGetVkGetInstanceProcAddr() unsafe.Pointer {
	return C.SDL_Vulkan_GetVkGetInstanceProcAddr()
}

// VulkanUnloadLibrary unloads the Vulkan loader library previously loaded by VulkanLoadLibrary().
// (https://wiki.libsdl.org/SDL_Vulkan_UnloadLibrary)
func VulkanUnloadLibrary() {
	C.SDL_Vulkan_UnloadLibrary()
}

// VulkanGetInstanceExtensions gets the names of the Vulkan instance extensions needed to create a surface with VulkanCreateSurface().
// (https://wiki.libsdl.org/SDL_Vulkan_GetInstanceExtensions)
func (window *Window) VulkanGetInstanceExtensions() []string {
	var count C.uint
	C.SDL_Vulkan_GetInstanceExtensions(window.cptr(), &count, nil)
	if count == 0 {
		return nil
	}

	strptrs := make([]*C.char, uint(count))
	C.SDL_Vulkan_GetInstanceExtensions(window.cptr(), &count, &strptrs[0])
	extensions := make([]string, uint(count))
	for i := range strptrs {
		extensions[i] = C.GoString(strptrs[i])
	}
	return extensions
}

// VulkanCreateSurface creates a Vulkan rendering surface for a window.
// (https://wiki.libsdl.org/SDL_Vulkan_CreateSurface)
func (window *Window) VulkanCreateSurface(instance interface{}) (surface unsafe.Pointer, err error) {
	if instance == nil {
		return nil, errors.New("vulkan: instance is nil")
	}
	val := reflect.ValueOf(instance)
	if val.Kind() != reflect.Ptr {
		return nil, errors.New("vulkan: instance is not a VkInstance (expected kind Ptr, got " + val.Kind().String() + ")")
	}
	var vulkanSurface C.VkSurfaceKHR
	if C.SDL_Vulkan_CreateSurface(window.cptr(),
		(C.VkInstance)(unsafe.Pointer(val.Pointer())),
		(*C.VkSurfaceKHR)(unsafe.Pointer(&vulkanSurface))) == C.SDL_FALSE {

		return nil, GetError()
	}
	return unsafe.Pointer(&vulkanSurface), nil
}

// VulkanGetDrawableSize gets the size of a window's underlying drawable in pixels (for use with setting viewport, scissor & etc).
// (https://wiki.libsdl.org/SDL_Vulkan_GetDrawableSize)
func (window *Window) VulkanGetDrawableSize() (w, h int32) {
	var _w, _h C.int
	C.SDL_Vulkan_GetDrawableSize(window.cptr(), &_w, &_h)
	return int32(_w), int32(_h)
}
