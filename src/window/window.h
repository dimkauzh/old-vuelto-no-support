#include <SDL2/SDL.h>
#include "stdio.h"
#include <stdbool.h>


SDL_Window* NewWindow(const char* title, int width, int height) {
	SDL_Window* window = SDL_CreateWindow(title,
                                       SDL_WINDOWPOS_CENTERED,
                                       SDL_WINDOWPOS_CENTERED,
                                       width, height, 0);
    if (window == NULL) {
        // Window creation failed
        printf("Window could not be created! SDL_Error: %s\n", SDL_GetError());
        SDL_Quit();
    } else {
        SDL_ShowWindow(window);
    }
    return window;
}

bool Loop(SDL_Window* window) {
    // Main loop flag
    bool quit = true;

    // Event handler
    SDL_Event e;

    // Main loop
    while (quit) {
        // Handle events on the queue
        while (SDL_PollEvent(&e)) {
            if (e.type == SDL_QUIT) {
                quit = false;
            }
        }
        return quit;
    }

    // Destroy window
    SDL_DestroyWindow(window);
    SDL_Quit();
    return quit;
}

void TestLoop(SDL_Window* window) {
    // Main loop flag
    bool quit = true;

    // Event handler
    SDL_Event e;

    // Main loop
    while (quit) {
        // Handle events on the queue
        while (SDL_PollEvent(&e)) {
            if (e.type == SDL_QUIT) {
                quit = false;
            }
        }
    }

    // Destroy window
    SDL_DestroyWindow(window);
    SDL_Quit();
}