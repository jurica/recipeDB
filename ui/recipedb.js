 $(document).ready(function() {
    $('#mainmenu').dropdown();

    $('#buttonHome').click(buttonHomeClick);
    $('#buttonNew').click(buttonNewClick);
    $('#buttonTags').click(buttonTagsClick);
    $('#buttonPrint').click(buttonPrintClick);

    buttonHomeClick();
});

function buttonNewClick() {
    $('#modalTest').modal('show');
}

function buttonHomeClick(){
    $('#modalLoading').modal('show');
    
    $.when(loadTemplate("recipeSimple"), loadRecipes()).done(function(template, recipes){
        console.log("ajax calls completed");
        console.log(template[0]);
        console.log(recipes[0]);
        output = Mustache.render(template[0], recipes[0]);
        console.log(output);
        $('#content').html(output);
        $('#modalLoading').modal('hide');
    });
}

function loadTemplate(name) {
    console.log("load template");
    return $.ajax({
        url: "templates/"+name+".tpl",
    });
}

function loadRecipes() {
    console.log("load recipes");
    return $.ajax({
        url: "recipe",
    });
}

function buttonTagsClick(){
    $.get("templates/recipe.tpl", function(recipeTpl){
        console.log(recipeTpl);
        $('#content').html(recipeTpl);
    });
}

function buttonPrintClick(){
}