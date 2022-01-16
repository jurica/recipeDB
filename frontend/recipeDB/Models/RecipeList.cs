using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;

namespace recipeDB.Models
{
    public class RecipeList
    {
        public int RecipeCount { get; set; }

        public int Offset { get; set; }

        public int Limit { get; set; }

        public int PageCount { get; set; }

        public int CurrentPage { get; set; }

        public Recipe[] Recipes { get; set; }

        public String SearchQuery { get; set; }
    }
}
