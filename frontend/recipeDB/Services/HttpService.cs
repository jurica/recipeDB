using recipeDB.Models;
using Microsoft.AspNetCore.Components;
using Microsoft.Extensions.Configuration;
using System;
using System.Collections.Generic;
using System.Net;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Net.Http.Json;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;

namespace recipeDB.Services
{
    public interface IHttpService
    {
        Task<T> Get<T>(string uri);
        Task<T> Post<T>(string uri, object value);
        Task<T> Put<T>(string uri, object value);
        Task<T> Delete<T>(string uri);
        void SetBaseAddress(Uri uri);
    }

    public class HttpService : IHttpService
    {
        private HttpClient _httpClient;
        private NavigationManager _navigationManager;
        private ILocalStorageService _localStorageService;

        public HttpService(NavigationManager navigationManager, ILocalStorageService localStorageService){
            _httpClient = new HttpClient();
            _navigationManager = navigationManager;
            _localStorageService = localStorageService;
        }

        public void SetBaseAddress(Uri uri) {
            _httpClient.BaseAddress = uri;
        }

        private  async Task<T> sendRequest<T>(HttpRequestMessage request) {
            var user = await _localStorageService.GetItem<User>("user");
            if (user != null) {
                request.Headers.Authorization = new System.Net.Http.Headers.AuthenticationHeaderValue("Bearer", user.Token);
            }

            using var response = await _httpClient.SendAsync(request);

            if (response.StatusCode == HttpStatusCode.Unauthorized) {
                var returnUrl = WebUtility.UrlEncode(new Uri(_navigationManager.Uri).PathAndQuery);
                _navigationManager.NavigateTo($"logout?returnUrl={returnUrl}");
                return default;
            }

            if (!response.IsSuccessStatusCode) {
                var error = await response.Content.ReadFromJsonAsync<Dictionary<string, string>>();
                throw new Exception(error["message"]);
            }

            return await response.Content.ReadFromJsonAsync<T>();
        }
        public async Task<T> Get<T>(string uri) {
            var request = new HttpRequestMessage(HttpMethod.Get, $"{_httpClient.BaseAddress}{uri}");
            return await sendRequest<T>(request);
        }

        public async Task<T> Post<T>(string uri, object value){
            var request = new HttpRequestMessage(HttpMethod.Post, $"{_httpClient.BaseAddress}{uri}");
            request.Content = new StringContent(JsonSerializer.Serialize(value), Encoding.UTF8, "application/json");
            return await sendRequest<T>(request);
        }

        public async Task<T> Put<T>(string uri, object value){
            var request = new HttpRequestMessage(HttpMethod.Put, $"{_httpClient.BaseAddress}{uri}");
            request.Content = new StringContent(JsonSerializer.Serialize(value), Encoding.UTF8, "application/json");
            return await sendRequest<T>(request);
        }

        public async Task<T> Delete<T>(string uri) {
            var request = new HttpRequestMessage(HttpMethod.Delete, $"{_httpClient.BaseAddress}{uri}");
            return await sendRequest<T>(request);
        }
    }
}