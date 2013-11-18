var async = require('async')
  , dataUtils = require('./data-utils') 
  ;

// Creates the 'Pipeline' object with the pipeline specific functions in it.
// Returns the async.compose results
var loadMgrDemog = (function() {
  var createUpdateObjects = function( pipeObj, callback ) {
    var docs = []
      , recs = pipeObj.data[ pipeObj.prevStep ]
      ;
   
    for(var i=0; i < recs.length; i++ ) {
      // check to see if managerID is empty. 
      // if so, don't do this
      if ( recs[i].managerID != '' && recs[i].managerID != null ) {
        var id = recs[i].managerID;
 
        delete recs[i].managerID;
        delete recs[i].playerID;
        delete recs[i].weight;
        delete recs[i].height;
        delete recs[i].bats;
        delete recs[i].throws;
        delete recs[i].debut;
        delete recs[i].finalGame;
        delete recs[i].college;
        delete recs[i].lahman40ID;
        delete recs[i].lahman45ID;
        delete recs[i].holtzID;
  
        docs.push( {query: {_id: id}, set: {$set: recs[i]} } );
      }
    }

    pipeObj.data.push( docs );
    pipeObj.prevStep += 1;
    callback( null, pipeObj );
  };


  return  async.compose(
     createUpdateObjects, 
     dataUtils.createObjects,
     dataUtils.readRemoteCsv
  );

}());


var inputSrc = { path: dataUtils.baseGithubUrl + '/Master.csv',
              headers: 1,
          dataTypeMap: [ 'managerID', 'bbrefID', 'birthCity', 'birthCountry', 
                         'birthState','deathCity', 'deathCountry', 'deathState', 'hofID', 'nameFirst', 'nameGiven', 'nameLast', 'nameNick',
                 	 'nameNote', 'retroID' ],
          floatColMap: [] };

var outputSettings = { type: 'mongodb', 
                        url: dataUtils.mongoUrl, 
                 collection: 'managers' };

var pipeObj = { input: inputSrc,
               exitFn: null,
               output: outputSettings,
                 data: [],
             prevStep: -1 };


loadMgrDemog( pipeObj, function(err, result ) {

  if ( pipeObj.output ) {
    dataUtils.updateFns[ pipeObj.output.type ]( result.data[2], pipeObj.output, function( err, r) {
      console.log('[loadMgrDemog] Manager Demographics loaded' );
    });
  }
});
