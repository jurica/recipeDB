using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;

namespace recipeDB.Models
{
    public class Recipe
    {
        public uint ID { get; set; }

        [Required]
        public String Name { get; set; }

        [ValidateComplexType]
        public List<Ingredient> Ingredients { get; set; }

        [ValidateComplexType]
        public List<Step> Steps { get; set; }

        public void addStepAfter(Step step)
        {
            int index = this.Steps.IndexOf(step) + 1;
            this.Steps.Insert(index, new Step());
        }

        public void removeStep(Step step)
        {
            this.Steps.Remove(step);
        }

        public void moveStepUp(Step step)
        {
            // TODO add nullcheck
            int indexToMove = this.Steps.IndexOf(step);
            Step tmp = step;
            this.Steps[indexToMove] = this.Steps[indexToMove - 1];
            this.Steps[indexToMove - 1] = tmp;
        }

        public void moveStepDown(Step step)
        {
            // TODO add nullcheck
            int indexToMove = this.Steps.IndexOf(step);
            Step tmp = step;
            this.Steps[indexToMove] = this.Steps[indexToMove + 1];
            this.Steps[indexToMove + 1] = tmp;
        }

        public void addIgredientAfter(Ingredient igredient)
        {
            int index = this.Ingredients.IndexOf(igredient) + 1;
            this.Ingredients.Insert(index, new Ingredient());
        }

        public void removeIngredient(Ingredient ingredient)
        {
            this.Ingredients.Remove(ingredient);
        }

        public void moveIngredientUp(Ingredient ingredient)
        {
            // TODO add nullcheck
            int indexToMove = this.Ingredients.IndexOf(ingredient);
            Ingredient tmp = ingredient;
            this.Ingredients[indexToMove] = this.Ingredients[indexToMove - 1];
            this.Ingredients[indexToMove - 1] = tmp;
        }

        public void moveIngredientDown(Ingredient ingredient)
        {
            // TODO add nullcheck
            int indexToMove = this.Ingredients.IndexOf(ingredient);
            Ingredient tmp = ingredient;
            this.Ingredients[indexToMove] = this.Ingredients[indexToMove + 1];
            this.Ingredients[indexToMove + 1] = tmp;
        }

        public string editUrl()
        {
            return "recipe/" + this.ID.ToString() + "/edit";
        }
    }
}
