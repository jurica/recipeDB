using recipeDB.Models;
using Microsoft.AspNetCore.Components;
using System.Threading.Tasks;
using System.Timers;
using recipeDB.Helpers;

namespace recipeDB.Services
{
    public interface IAuthenticationService
    {
        User User { get; }
        Task Initialize();
        Task Login(User user);
        Task Logout();
    }

    public class AuthenticationService : IAuthenticationService
    {
        private IHttpService _httpService;
        private NavigationManager _navigationManager;
        private ILocalStorageService _localStorageService;
        private Timer _tokenRefreshTimer;

        public User User { get; private set; }

        public AuthenticationService(
            IHttpService httpService,
            NavigationManager navigationManager,
            ILocalStorageService localStorageService
        ) {
            _httpService = httpService;
            _navigationManager = navigationManager;
            _localStorageService = localStorageService;

            _tokenRefreshTimer = new Timer(5 * 60 * 1000);
            _tokenRefreshTimer.Elapsed += refreshToken;
            _tokenRefreshTimer.AutoReset = true;
            _tokenRefreshTimer.Enabled = false;
        }

        private async void refreshToken(object source, ElapsedEventArgs e) {
            if (User != null) {
                User = await _httpService.Post<User>("/refresh-token", null);
                await _localStorageService.SetItem("user", User);
            } else {
                _tokenRefreshTimer.Stop();
            }
        }
        public async Task Initialize()
        {
            User = await _localStorageService.GetItem<User>("user");
            if (User != null) {
                _tokenRefreshTimer.Start();
            }
        }

        public async Task Login(User user)
        {
            User = await _httpService.Post<User>("login", user);
            await _localStorageService.SetItem("user", User);
            _tokenRefreshTimer.Start();
        }

        public async Task Logout()
        {
            _tokenRefreshTimer.Stop();
            var returnUrl = _navigationManager.QueryString("returnUrl") ?? "/";
            User = null;
            await _localStorageService.RemoveItem("user");
            _navigationManager.NavigateTo($"login?returnUrl={returnUrl}");
        }
    }
}