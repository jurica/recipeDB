using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;

namespace recipeDB
{
    public class RecipeModel
    {
        public uint ID { get; set; }

        [Required]
        public String Name { get; set; }

        [ValidateComplexType]
        public List<IngredientModel> Ingredients { get; set; }

        [ValidateComplexType]
        public List<StepModel> Steps { get; set; }

        public void addStepAfter(StepModel step)
        {
            int index = this.Steps.IndexOf(step) + 1;
            this.Steps.Insert(index, new StepModel());
        }

        public void removeStep(StepModel step)
        {
            this.Steps.Remove(step);
        }

        public void moveStepUp(StepModel step)
        {
            // TODO add nullcheck
            int indexToMove = this.Steps.IndexOf(step);
            StepModel tmp = step;
            this.Steps[indexToMove] = this.Steps[indexToMove - 1];
            this.Steps[indexToMove - 1] = tmp;
        }

        public void moveStepDown(StepModel step)
        {
            // TODO add nullcheck
            int indexToMove = this.Steps.IndexOf(step);
            StepModel tmp = step;
            this.Steps[indexToMove] = this.Steps[indexToMove + 1];
            this.Steps[indexToMove + 1] = tmp;
        }

        public void addIgredientAfter(IngredientModel igredient)
        {
            int index = this.Ingredients.IndexOf(igredient) + 1;
            this.Ingredients.Insert(index, new IngredientModel());
        }

        public void removeIngredient(IngredientModel ingredient)
        {
            this.Ingredients.Remove(ingredient);
        }

        public void moveIngredientUp(IngredientModel ingredient)
        {
            // TODO add nullcheck
            int indexToMove = this.Ingredients.IndexOf(ingredient);
            IngredientModel tmp = ingredient;
            this.Ingredients[indexToMove] = this.Ingredients[indexToMove - 1];
            this.Ingredients[indexToMove - 1] = tmp;
        }

        public void moveIngredientDown(IngredientModel ingredient)
        {
            // TODO add nullcheck
            int indexToMove = this.Ingredients.IndexOf(ingredient);
            IngredientModel tmp = ingredient;
            this.Ingredients[indexToMove] = this.Ingredients[indexToMove + 1];
            this.Ingredients[indexToMove + 1] = tmp;
        }

        public string editUrl()
        {
            return "recipe/" + this.ID.ToString() + "/edit";
        }
    }
}
