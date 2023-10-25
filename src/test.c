#include "stdio.h"
#include <SDL2/SDL.h>

void window_test() {
    // Initialize SDL
    if (SDL_Init(SDL_INIT_VIDEO) < 0) {
        // Initialization failed
        printf("SDL could not initialize! SDL_Error: %s\n", SDL_GetError());
    }

    // Create a window
    SDL_Window* window = SDL_CreateWindow("SDL2 Window", SDL_WINDOWPOS_UNDEFINED, SDL_WINDOWPOS_UNDEFINED, 800, 600, SDL_WINDOW_SHOWN);
    if (window == NULL) {
        // Window creation failed
        printf("Window could not be created! SDL_Error: %s\n", SDL_GetError());
        SDL_Quit();
    }
    SDL_Renderer* renderer = SDL_CreateRenderer(window, -1, SDL_RENDERER_ACCELERATED);

	if (renderer == NULL) {
    	// Renderer creation failed
    	printf("Renderer could not be created! SDL_Error: %s\n", SDL_GetError());
    	SDL_DestroyWindow(window);
    	SDL_Quit();
	}

    // Main loop flag
    int quit = 0;

    // Event handler
    SDL_Event e;

    // Main loop
    while (!quit) {
        // Handle events on the queue
        while (SDL_PollEvent(&e) != 0) {
            if (e.type == SDL_QUIT) {
                quit = 1;
            }
        }

        // Clear the screen (you can set your own color)
        SDL_SetRenderDrawColor(renderer, 0, 0, 0, 255);
        SDL_RenderClear(renderer);

        // Update the screen
        SDL_RenderPresent(renderer);
    }

    // Destroy the window
    SDL_DestroyWindow(window);

    // Quit SDL
    SDL_Quit();

}

void test() {
	printf("test");
}
