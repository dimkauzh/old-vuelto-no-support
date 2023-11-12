#include <SDL2/SDL.h>
#include <stdio.h>
#include <stdbool.h>

int SCREEN_WIDTH;
int SCREEN_HEIGHT;

int FPS = 60;
int FPS_ON = 0;

SDL_Texture *pixelTexture;
Uint32 *pixelBuffer;

SDL_Window *window;
SDL_Renderer *renderer;

SDL_Event event;


SDL_Window *NewWindow(const char *title, int width, int height)
{

    SCREEN_WIDTH = width;
    SCREEN_HEIGHT = height;

    SDL_Window *win = SDL_CreateWindow(title, SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED, width, height, 0);
    if (win == NULL)
    {
        // Window creation failed
        printf("Window could not be created! SDL_Error: %s\n", SDL_GetError());
        SDL_Quit();
    }
    else
    {
        SDL_ShowWindow(win);
    }

    window = win;

    renderer = SDL_CreateRenderer(win, -1, SDL_RENDERER_SOFTWARE);
    if (renderer == NULL)
    {
        // Renderer creation failed
        printf("Renderer could not be created! SDL_Error: %s\n", SDL_GetError());
        SDL_DestroyWindow(win);
        SDL_Quit();
    }

    pixelTexture = SDL_CreateTexture(renderer, SDL_PIXELFORMAT_RGBA8888, SDL_TEXTUREACCESS_STREAMING, SCREEN_WIDTH, SCREEN_HEIGHT);
    pixelBuffer = (Uint32 *)malloc(SCREEN_WIDTH * SCREEN_HEIGHT * sizeof(Uint32));

    return win;
}

void SetIcon(const char *iconPath) {
    SDL_Surface* iconSurface = SDL_LoadBMP(iconPath);
    if (iconSurface) {
        SDL_SetWindowIcon(window, iconSurface);
        SDL_FreeSurface(iconSurface);
    } else {
        printf("Failed to load icon: %s\n", SDL_GetError());
    }
}

void ResetScreen()
{
    Uint32 clearColor = SDL_MapRGB(SDL_AllocFormat(SDL_PIXELFORMAT_RGBA8888), 0, 0, 0);

    for (int i = 0; i < SCREEN_WIDTH * SCREEN_HEIGHT; i++)
    {
        pixelBuffer[i] = clearColor;
    }
}

void SetFPS(int fps)
{
    FPS = fps;
}

int GetFPS()
{
    return FPS;
}

void SetDelay(int delay)
{
    SDL_Delay(delay);
}

void RemoveFPSLimit()
{
    FPS_ON = 1;
}

void LimitFrameRate(Uint32 frameStartTime) {
    Uint32 frameTime = SDL_GetTicks() - frameStartTime;
    if (frameTime < 1000 / FPS) {
        SDL_Delay((1000 / FPS) - frameTime);
    }
}

void Update(Uint32 frameStartTime)
{
    SDL_UpdateTexture(pixelTexture, NULL, pixelBuffer, SCREEN_WIDTH * sizeof(Uint32));
    SDL_RenderCopy(renderer, pixelTexture, NULL, NULL);
    SDL_RenderPresent(renderer);

    if (FPS_ON == 0) {
        LimitFrameRate(frameStartTime);
    }

    ResetScreen();
}

void Quit()
{
    free(pixelBuffer);

    SDL_DestroyTexture(pixelTexture);
    SDL_DestroyRenderer(renderer);
    SDL_DestroyWindow(window);

    SDL_Quit();
}

bool Loop()
{
    bool running = true;

    while (running)
    {
        Uint32 frameStartTime = SDL_GetTicks();

        while (SDL_PollEvent(&event))
        {
            if (event.type == SDL_QUIT)
            {
                running = false;
            }
        }

        Update(frameStartTime);

        return running;
    }

    Quit();

    return running;
}

void SetPixel(int x, int y, int r, int g, int b)
{
    if (x >= 0 && x < SCREEN_WIDTH && y >= 0 && y < SCREEN_HEIGHT)
    {
        pixelBuffer[y * SCREEN_WIDTH + x] = SDL_MapRGB(SDL_AllocFormat(SDL_PIXELFORMAT_RGBA8888), r, g, b);
    }
}

void DrawRect(int x, int y, int width, int height, int r, int g, int b)
{
    Uint32 color = SDL_MapRGB(SDL_AllocFormat(SDL_PIXELFORMAT_RGBA8888), r, g, b);
    Uint32* dst = pixelBuffer + y * SCREEN_WIDTH + x;

    for (int i = 0; i < height; i++) {
        Uint32* row = dst;
        for (int j = 0; j < width; j++) {
            row[j] = color;
        }
        dst += SCREEN_WIDTH;
    }
}

void SetBackgroundColor(int r, int g, int b)
{
    Uint32 clearColor = SDL_MapRGB(SDL_AllocFormat(SDL_PIXELFORMAT_RGBA8888), r, g, b);

    for (int i = 0; i < SCREEN_WIDTH * SCREEN_HEIGHT; i++)
    {
        pixelBuffer[i] = clearColor;
    }
}

// Function to handle events for the button
bool IsButtonPressed(int x, int y, int width, int height) {
    if (SDL_PollEvent(&event)) {
        if (event.type == SDL_MOUSEBUTTONDOWN) {
            int mouseX, mouseY;
            SDL_GetMouseState(&mouseX, &mouseY);

            if (mouseX >= x && mouseX <= x + width &&
                mouseY >= y && mouseY <= y + height) {
                return true;
            }
            return false;
        }
        return false;
    }

    return false;
}

