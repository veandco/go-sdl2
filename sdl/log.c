#include "_cgo_export.h"

void LogSetOutputFunction(void *data)
{
    SDL_LogSetOutputFunction((SDL_LogOutputFunction)logOutputFunction, data);
}
