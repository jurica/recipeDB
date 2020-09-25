using System;
namespace recipeDB
{
    public class IngredientModel
    {
        public string Name { get; set; }
        public string Amount { get; set; }
        public string Unit { get; set; }
        public int SortOrder { get; set; }
}
}
