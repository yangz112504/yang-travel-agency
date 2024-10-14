const searchBtn = document.getElementById("searchButton");
const clearBtn = document.getElementById("clearButton");
const resultDiv = document.getElementById("results");

function clearResults(){
    resultDiv.innerHTML= "";
}

function generateResult(name, imageUrl, description){
    resultDiv.innerHTML+=`
    <div class="card results-card">
        <img src="${imageUrl}" class="card-img-top">
        <div class="card-body">
            <h5 class="card-title">${name}</h5>
            <p class="card-text">
                ${description}
            </p>
            <a href="#" class="btn btn-info">Visit</a>
        </div>
    </div>
    `;    
}
function search(){
    const input = document.getElementById("search-bar").value.toLowerCase();
    console.log(input);
    clearResults();
    fetch('./travel_recommendation_api.json')
        .then(response => response.json())
        .then(data=>{
            if(input === "countries" || input === "country"){
                data.countries.forEach(function(country){
                    country.cities.forEach(function(city){
                        generateResult(city.name, city.imageUrl, city.description);
                    })
                })
            }
            if(input === "temple" || input === "temples"){
                data.temples.forEach(function(temple){
                    generateResult(temple.name, temple.imageUrl, temple.description);
                })
            }
            if(input === "beach" || input === "beaches"){
                data.beaches.forEach(function(beach){
                    generateResult(beach.name, beach.imageUrl, beach.description);
                })
            }
        })
        .catch(error =>{
            resultDiv.innerHTML+=`
                <div class="card results-card">
                    <div class="card-body">
                        <h5 class="card-title">An error occurred while fetching data.</h5>
                    </div>
                 </div>
                 `;    
        });
}

function thankyou(){
    alert('Thank you for contacting us! We received your message.');
}

searchBtn.addEventListener("click", search)
clearBtn.addEventListener("click",clearResults);