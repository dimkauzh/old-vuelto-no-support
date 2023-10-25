#include <SDL2/SDL.h>
#include "stdio.h"
#include <stdbool.h>


SDL_Window* NewWindow(const char* title, int width, int height) {
	SDL_Window* window = SDL_CreateWindow("SDL2 Window", SDL_WINDOWPOS_UNDEFINED, SDL_WINDOWPOS_UNDEFINED, 800, 600, SDL_WINDOW_SHOWN);
    if (window == NULL) {
        // Window creation failed
        printf("Window could not be created! SDL_Error: %s\n", SDL_GetError());
        SDL_Quit();
    }
    return window;
}

bool Loop(SDL_Window* window) {
    // Main loop flag
    bool quit = false;

    // Event handler
    SDL_Event e;

    // Main loop
    while (!quit) {
        // Handle events on the queue
        while (SDL_PollEvent(&e) != 0) {
            if (e.type == SDL_QUIT) {
                quit = true;
                SDL_DestroyWindow(window);
				SDL_Quit();
            }
        }
    }

    return quit;
}