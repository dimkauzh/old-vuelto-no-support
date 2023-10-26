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

SDL_Renderer* NewRenderer(SDL_Window* window) {
    SDL_Renderer* renderer = SDL_CreateRenderer(window, -1, SDL_RENDERER_SOFTWARE);
    if (renderer == NULL) {
        // Window creation failed
        printf("Renderer could not be created! SDL_Error: %s\n", SDL_GetError());
        SDL_DestroyWindow(window);
        SDL_Quit();
    }
    return renderer;
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