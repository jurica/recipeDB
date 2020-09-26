using System;
using System.ComponentModel.DataAnnotations;

namespace recipeDB
{
    public class IngredientModel
    {
        [Required]
        public string Name { get; set; }
        
        [Required]
        public string Amount { get; set; }
        
        [Required]
        public string Unit { get; set; }
}
}
