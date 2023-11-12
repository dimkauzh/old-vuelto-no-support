#ifdef WINDOW_H
#include "window.c"
#endif

#include <SDL2/SDL.h>
#include <stdbool.h>



SDL_Window *NewWindow(const char *title, int width, int height);
void SetIcon(const char *iconPath);
void Update();
void Quit();
bool Loop();

void SetPixel(int x, int y, int r, int g, int b);
void DrawRect(int x, int y, int width, int height, int r, int g, int b);
void SetBackgroundColor(int r, int g, int b);
void ResetScreen();

void SetFPS(int fps);
int GetFPS();
void RemoveFPSLimit();
void SetDelay(int delay);

bool IsButtonPressed(int x, int y, int width, int height);