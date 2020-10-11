using System;
using System.ComponentModel.DataAnnotations;

namespace recipeDB.Models
{
    public class Step
    {
        public uint RecipeID { get; set; }
        
        [Required]
        public String Description { get; set; }
    }
}