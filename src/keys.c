#include <SDL2/SDL.h>
#include <stdbool.h>

bool IsKeyPressed(SDL_Scancode key)
{
    const Uint8 *state = SDL_GetKeyboardState(NULL);

    if (state[key])
    {
        return true;
    }

    return false;
}