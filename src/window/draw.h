#include <SDL2/SDL.h>

void SetPixel(SDL_Renderer* renderer, int x, int y, int r, int g, int b) {
    SDL_SetRenderDrawColor(renderer, r, g, b, 255);
    SDL_RenderDrawPoint(renderer, x, y);
}

void DrawRect(SDL_Renderer* renderer, int x, int y, int width, int height, int r, int g, int b) {
    for (int i = x; i < x + width; i++) {
        for (int j = y; j < y + height; j++) {
            SetPixel(renderer, i, j, r, g, b);
        }
    }
}
