#include <SDL2/SDL.h>
#include <stdio.h>
#include <stdbool.h>

int SCREEN_WIDTH;
int SCREEN_HEIGHT;

int FPS = 60;

SDL_Texture* pixelTexture;
Uint32* pixelBuffer;

SDL_Window* window;
SDL_Renderer* renderer;

SDL_Window* NewWindow(const char* title, int width, int height) {
    
    SCREEN_WIDTH = width;
    SCREEN_HEIGHT = height;

    SDL_Window* win = SDL_CreateWindow(title, SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED, width, height, 0);
    if (win == NULL) {
        // Window creation failed
        printf("Window could not be created! SDL_Error: %s\n", SDL_GetError());
        SDL_Quit();
    } else {
        SDL_ShowWindow(win);
    }

    renderer = SDL_CreateRenderer(win, -1, SDL_RENDERER_SOFTWARE);
    if (renderer == NULL) {
        // Renderer creation failed
        printf("Renderer could not be created! SDL_Error: %s\n", SDL_GetError());
        SDL_DestroyWindow(win);
        SDL_Quit();
    }

    pixelTexture = SDL_CreateTexture(renderer, SDL_PIXELFORMAT_RGBA8888, SDL_TEXTUREACCESS_STREAMING, SCREEN_WIDTH, SCREEN_HEIGHT);
    pixelBuffer = (Uint32*)malloc(SCREEN_WIDTH * SCREEN_HEIGHT * sizeof(Uint32));

    return win;
}

void ResetScreen() {
    Uint32 clearColor = SDL_MapRGB(SDL_AllocFormat(SDL_PIXELFORMAT_RGBA8888), 0, 0, 0);

    for (int i = 0; i < SCREEN_WIDTH * SCREEN_HEIGHT; i++) {
        pixelBuffer[i] = clearColor;
    }
}

void SetFPS(int fps) {
    FPS = fps;
}

int GetFPS() {
    return FPS;
}


void Update() {
    SDL_UpdateTexture(pixelTexture, NULL, pixelBuffer, SCREEN_WIDTH * sizeof(Uint32));
    SDL_RenderCopy(renderer, pixelTexture, NULL, NULL);
    SDL_RenderPresent(renderer);

    SDL_Delay(1000 / FPS);



    ResetScreen();
}

void Quit() {
    free(pixelBuffer);

    SDL_DestroyTexture(pixelTexture);
    SDL_DestroyRenderer(renderer);
    SDL_DestroyWindow(window);

    SDL_Quit();
}

bool Loop() {
    bool running = true;

    SDL_Event e;

    while (running) {
        while (SDL_PollEvent(&e)) {
            if (e.type == SDL_QUIT) {
                running = false;
            }
        }

        Update();
        return running;
    }

    Quit();

    return running;
}

void SetPixel(int x, int y, int r, int g, int b) {
    if (x >= 0 && x < SCREEN_WIDTH && y >= 0 && y < SCREEN_HEIGHT) {
        pixelBuffer[y * SCREEN_WIDTH + x] = SDL_MapRGB(SDL_AllocFormat(SDL_PIXELFORMAT_RGBA8888), r, g, b);
    }
}

void DrawRect(int x, int y, int width, int height, int r, int g, int b) {
    for (int i = x; i < x + width; i++) {
        for (int j = y; j < y + height; j++) {
            SetPixel(i, j, r, g, b);
        }
    }
}

void SetBackgroundColor(int r, int g, int b) {
    Uint32 clearColor = SDL_MapRGB(SDL_AllocFormat(SDL_PIXELFORMAT_RGBA8888), r, g, b);

    for (int i = 0; i < SCREEN_WIDTH * SCREEN_HEIGHT; i++) {
        pixelBuffer[i] = clearColor;
    }
}
