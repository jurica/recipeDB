using System;
using System.Net.Http;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Components.WebAssembly.Hosting;
using Microsoft.Extensions.DependencyInjection;
using recipeDB.Services;

namespace recipeDB
{
    public class Program
    {
        public static async Task Main(string[] args)
        {
            var builder = WebAssemblyHostBuilder.CreateDefault(args);
            builder.RootComponents.Add<App>("app");

            builder.Services
            .AddScoped<IAuthenticationService, AuthenticationService>()
            .AddScoped<ILocalStorageService, LocalStorageService>()
            .AddScoped<IHttpService, HttpService>();

            var host = builder.Build();

            var authenticationService = host.Services.GetRequiredService<IAuthenticationService>();
            await authenticationService.Initialize();

            var httpService = host.Services.GetRequiredService<IHttpService>();
            httpService.SetBaseAddress(new Uri(builder.Configuration["apiUrl"]));

            await host.RunAsync();
        }
    }
}
