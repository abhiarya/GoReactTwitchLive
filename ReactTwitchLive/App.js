import React from 'react';
import { AppRegistry, TouchableOpacity, StyleSheet,Alert, Linking, ScrollView, Text, View, Image, Button, YellowBox } from 'react-native';

// This is to suppress the warining of some deprecated functions used by React native code
YellowBox.ignoreWarnings([
  'Warning: componentWillMount is deprecated',
  'Warning: componentWillReceiveProps is deprecated',
]);

//Function component that holds one live stream result
const LiveStream = (props) => {
    return (
    <View style={styles.style_liveStream}>
      <Text style={styles.style_UserName}> {props.UserName}</Text> 
      <TouchableOpacity onPress={() => Linking.openURL(props.LiveURL)}>
        <Image source={{uri: props.PreviewImage}}  style={{width: 320, height: 180}} />
      </TouchableOpacity>
    </View> 
  );
};

// Function component that has the list of all the liveStreams
const LiveStreamList = (props) => {
  return(
  	<ScrollView>
			{props.liveStreams.map(liveStream => <LiveStream key={liveStream.ID}  {...liveStream}/>)}
    </ScrollView>
  );
};

export default class App extends React.Component {

  state = {
    liveStreams: [],
    error: null,
    buttonLabel: 'Fetch Live Streams',
    welcomeText:  'Welcome!',
    //showLiveStreams: false,
    //user: '',
  };

  // login = () => {
  //   //Call the twich API to login

  //   fetch('http://localhost:8080/authorize')
  //     .then(response => response.json())
  //     .then(data => {
  //       this.setState(
  //         {user:data.Name} // Set the JSON response obtained to the liveStreams state
  //       );
  //       this.setState({showLiveStreams: true});
  //       this.setState({welcomeText: 'Welcome '+ {user} });
  //     })
  //     .catch(error => {
  //       this.setState({ error: error.message });
  //       Alert.alert(this.state.error);
  //     });
  // };


  fetchLiveStreams = () => {

    this.setState({liveStreams: []});

    fetch('http://localhost:8080/go-twitch-live')
      .then(response => response.json())
      .then(liveStreams => {
         this.setState(
          {liveStreams} // Set the JSON response obtained to the liveStreams state
        );
        this.setState({buttonLabel: 'Refresh'}); // Change the label of the button
      })
      .catch(error => {
        this.setState({ error: error.message });
        Alert.alert(this.state.error);
      });
  };

  render() {
    return (
      <View style={styles.container}>
        <Text style={styles.style_title} >Twitch Live</Text>
      
        <Text style= {styles.style_welcome}>{this.state.welcomeText}</Text>
      
        <TouchableOpacity
          style={styles.style_button}
          onPress={this.fetchLiveStreams}
          >
          buttonName: 'Fetch Live Streams',
          <Text style={{color: 'white', fontSize:18}} > {this.state.buttonLabel}</Text>
        </TouchableOpacity>

        <LiveStreamList liveStreams={this.state.liveStreams}/>       
      </View>
    );
  }
}


const styles = StyleSheet.create({
  container: {
    //flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },

  style_liveStream: {
    //flex: 1,
    backgroundColor: '#D3D3D3',
    alignItems: 'center',
    justifyContent: 'center',
    margin: 10,
  },

  style_title: {
    color: '#6441A5',
    fontWeight: 'bold',
    fontSize: 30,
    margin: 30,
  },

  style_welcome: {
    color: 'black',
    fontSize: 15,
    margin: 10,
  },


  style_UserName: {
    color: 'black',
    fontWeight: 'bold',
    fontSize: 18,
  },

  style_liveLink: {
    color: 'blue',
    fontSize: 17,
    textDecorationLine: 'underline',
  },

  style_buttonView: {
    margin: 20,
  },

  style_button: {
    alignItems: 'center',
    backgroundColor: '#0085FF',
    padding: 10,
    margin:10,
  },

  // style_login: {
  //   alignItems: 'center',
  //   backgroundColor: '#6441A5',
  //   padding: 10,
  //   margin:10,
  // },

});
