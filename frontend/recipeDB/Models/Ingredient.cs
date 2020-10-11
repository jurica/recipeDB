using System;
using System.ComponentModel.DataAnnotations;

namespace recipeDB.Models
{
    public class Ingredient
    {
        public uint RecipeID { get; set; }
        
        [Required]
        public string Name { get; set; }
        
        [Required]
        public string Amount { get; set; }
        
        [Required]
        public string Unit { get; set; }
}
}
