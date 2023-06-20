
# How to Save the Graph

This allows you to save the current graph of a Foam repo to a file:

1. Open the VSCode Developer Tools
2. Navigate to the correct iframe
3. Set a breakpoint in `graph.js` in the highlight update function (currently line $159$ -- the model will be object `m`)
4. Highlight a different node to trigger the breakpoint
5. Enter the following in the console (from [here](https://gist.github.com/raecoo/dcbac9e94198dfd0801be8a0cbb14570#file-console-save-js-L4))
   ```js
   // e.g. console.save({hello: 'world'})
   (function(console){
     console.save = function(data, filename){
       if(!data) {
         console.error('Console.save: No data')
         return;
       }
       if(!filename) filename = 'console.json'
       if(typeof data === "object"){
         data = JSON.stringify(data, undefined, 4)
       }
       var blob = new Blob([data], {type: 'text/json'}),
           e    = document.createEvent('MouseEvents'),
           a    = document.createElement('a')
       a.download = filename
       a.href = window.URL.createObjectURL(blob)
       a.dataset.downloadurl =  ['text/json', a.download, a.href].join(':')
       e.initMouseEvent('click', true, false, window, 0, 0, 0, 0, 0, false, false, false, false, 0, null)
       a.dispatchEvent(e)
     }
   })(console)
   ```
5. Now enter:
   ```js
   console.save(m)
   ```
   adjust the name of the model object if you set the breakpoint somewhere else.

