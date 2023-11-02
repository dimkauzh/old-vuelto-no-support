#ifdef WINDOW_H
#include "window.c"
#endif

#include <SDL2/SDL.h>
#include <stdbool.h>

SDL_Window* NewWindow(const char* title, int width, int height);
void Update();
void Quit();
bool Loop();

void SetPixel(int x, int y, int r, int g, int b);
void DrawRect(int x, int y, int width, int height, int r, int g, int b);
void SetBackgroundColor(int r, int g, int b);

void SetFPS(int fps);
int GetFPS();