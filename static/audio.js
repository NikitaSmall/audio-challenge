$(document).ready(function() {
  /* Copyright 2013 Chris Wilson

     Licensed under the Apache License, Version 2.0 (the "License");
     you may not use this file except in compliance with the License.
     You may obtain a copy of the License at

         http://www.apache.org/licenses/LICENSE-2.0

     Unless required by applicable law or agreed to in writing, software
     distributed under the License is distributed on an "AS IS" BASIS,
     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
     See the License for the specific language governing permissions and
     limitations under the License.
  */

  window.AudioContext = window.AudioContext || window.webkitAudioContext;

  var audioContext = new AudioContext();
  var audioInput = null,
      realAudioInput = null,
      inputPoint = null,
      audioRecorder = null;
  var rafID = null;
  var analyserContext = null;
  var recIndex = 0;

  /* TODO:

  - offer mono option
  - "Monitor input" switch
  */

  function saveAudio() {
      audioRecorder.exportWAV( doneEncoding );
      // could get mono instead by saying
      // audioRecorder.exportMonoWAV( doneEncoding );
  }

  function gotBuffers( buffers ) {
      // the ONLY time gotBuffers is called is right after a new recording is completed -
      // so here's where we should set up the download.
      audioRecorder.exportWAV( doneEncoding );
  }

  function doneEncoding( blob ) {
      console.log(blob);

      var data = new FormData();
      data.append('blob', blob);

      $.ajax({
        method: 'POST',
        url: '/message',
        data: data,
        contentType: false,
        processData: false
      }).done(function(data) {
        // console.log(data);
      });
      // Recorder.setupDownload( blob, "myRecording" + ((recIndex<10)?"0":"") + recIndex + ".wav" );
      recIndex++;
  }

  $('#save').click(function(e) {
    e.preventDefault();

    audioRecorder.exportWAV(function(blob) {

    });
  });

  $('#recorder').click(function() {
    if ($(this).hasClass("recording")) {
        // stop recording
        audioRecorder.stop();
        $(this).toggleClass("recording");
        audioRecorder.getBuffers( gotBuffers );
    } else {
      // start recording
      if (!audioRecorder)
          return;
      $(this).toggleClass("recording");
      audioRecorder.clear();
      audioRecorder.record();
    }
  });

  function convertToMono( input ) {
      var splitter = audioContext.createChannelSplitter(2);
      var merger = audioContext.createChannelMerger(2);

      input.connect( splitter );
      splitter.connect( merger, 0, 0 );
      splitter.connect( merger, 0, 1 );
      return merger;
  }

  function cancelAnalyserUpdates() {
      window.cancelAnimationFrame( rafID );
      rafID = null;
  }

  function toggleMono() {
      if (audioInput != realAudioInput) {
          audioInput.disconnect();
          realAudioInput.disconnect();
          audioInput = realAudioInput;
      } else {
          realAudioInput.disconnect();
          audioInput = convertToMono( realAudioInput );
      }

      audioInput.connect(inputPoint);
  }

  function gotStream(stream) {
      inputPoint = audioContext.createGain();

      // Create an AudioNode from the stream.
      realAudioInput = audioContext.createMediaStreamSource(stream);
      audioInput = realAudioInput;
      audioInput.connect(inputPoint);

  //    audioInput = convertToMono( input );

      analyserNode = audioContext.createAnalyser();
      analyserNode.fftSize = 2048;
      inputPoint.connect( analyserNode );

      audioRecorder = new Recorder( inputPoint );

      zeroGain = audioContext.createGain();
      zeroGain.gain.value = 0.0;
      inputPoint.connect( zeroGain );
      zeroGain.connect( audioContext.destination );
  }

  function initAudio() {
          if (!navigator.getUserMedia)
              navigator.getUserMedia = navigator.webkitGetUserMedia || navigator.mozGetUserMedia;
          if (!navigator.cancelAnimationFrame)
              navigator.cancelAnimationFrame = navigator.webkitCancelAnimationFrame || navigator.mozCancelAnimationFrame;
          if (!navigator.requestAnimationFrame)
              navigator.requestAnimationFrame = navigator.webkitRequestAnimationFrame || navigator.mozRequestAnimationFrame;

      navigator.getUserMedia(
          {
              "audio": {
                  "mandatory": {
                      "googEchoCancellation": "false",
                      "googAutoGainControl": "false",
                      "googNoiseSuppression": "false",
                      "googHighpassFilter": "false"
                  },
                  "optional": []
              },
          }, gotStream, function(e) {
              alert('Error getting audio');
              console.log(e);
          });
  }

  window.addEventListener('load', initAudio );
});
