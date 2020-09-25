using System;
using System.Collections.Generic;

namespace recipeDB
{
    public class RecipeModel
    {
        public Guid ID { get; set; }
        public String Name { get; set; }
        public List<IngredientModel> Ingredients { get; set; }
        public List<StepModel> Steps { get; set; }
    }

    public class StepModel
    {
        public String Description { get; set; }
    }
}
