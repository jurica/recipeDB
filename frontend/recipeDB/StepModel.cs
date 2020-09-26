using System;
using System.ComponentModel.DataAnnotations;

namespace recipeDB
{
    public class StepModel
    {
        [Required]
        public String Description { get; set; }
    }
}