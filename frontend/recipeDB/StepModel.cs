using System;
using System.ComponentModel.DataAnnotations;

namespace recipeDB
{
    public class StepModel
    {
        public uint RecipeID { get; set; }
        
        [Required]
        public String Description { get; set; }
    }
}