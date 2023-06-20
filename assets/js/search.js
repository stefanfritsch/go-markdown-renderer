const query_search = async (search_term) => {
  
  const query_body = {"term": search_term};
  
  const response = 
    await fetch('/api/search', {
      method: 'POST',
      body: JSON.stringify(query_body),
      headers: {
        "Accept": "application/json",
        "Content-Type": "application/json"
      }
    });
  const results = await response.json();
  
  return results
}

const render_search_results = (results) => {
  
  if (results.length === undefined && results.hasOwnProperty("message"))
    return `<div class="search-results-no-result">${results.message}</div>`
  if (results.length === 0) 
    return `<div class="search-results-no-result">No results</div>`
  
  var output = "";
  for (const result of results) 
  {
    output += render_search_result(result);
  }
  
  return output
}

const render_search_result = (result) => {
  var samples="";
  
  if (result.sample !== null)
  {
    for (const s of result.sample) 
    {
      samples = samples + `<div class="search-result-sample">${s}</div>`;
    }
  }

  return `<li class = "search-result">
    <a href="${result.path}">
      <div class="search-result-title">${result.name}</div>
      <div class="search-result-samples">${samples}</div>
    </a>
  </li>`
}

const search = async (search_term) => {
  const node = document.getElementById("search-results");
  if (search_term.length == 0) 
  {
    node.textContent = "";
  } else {
    const results = await query_search(search_term);
    const html_results = render_search_results(results);
    
    node.innerHTML = html_results;
  }
}

const clear_search_results = async () => {
  const results = document.getElementById("search-results");
  results.textContent = "";
  const search_query = document.getElementById("search-query");
  search_query.value = "";
}