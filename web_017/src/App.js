import React from 'react';
import CardContainer from './ProductCards';
import { BrowserRouter as Router, Route } from "react-router-dom";
import Nav from './Navigation';
import cookie from 'js-cookie';


class App extends React.Component {

  constructor(props) {
    super(props);
    const user = cookie.getJSON("user") || {loggedin:true};
    this.state = {
      user: user,
      showSignInModal: false,
      showBuyModal: false
    };
    this.showSignInModalWindow = this.showSignInModalWindow.bind(this);
    this.toggleSignInModalWindow = this.toggleSignInModalWindow.bind(this);
    this.showBuyModalWindow = this.showBuyModalWindow.bind(this);
    this.toggleBuyModalWindow = this.toggleBuyModalWindow.bind(this);
    
  }

  showSignInModalWindow(){
    const state = this.state;
    const newState = Object.assign({},state,{showSignInModal:true});
    this.setState(newState);
  }

  toggleSignInModalWindow() {
    const state = this.state;
    const newState = Object.assign({},state,{showSignInModal:!state.showSignInModal});
    this.setState(newState);
  }



  showBuyModalWindow(id,price){
    const state = this.state;
    const newState = Object.assign({},state,{showBuyModal:true,productid:id,price:price});
    this.setState(newState);
  }

  toggleBuyModalWindow(){
    const state = this.state;
    const newState = Object.assign({},state,{showBuyModal:!state.showBuyModal});
    this.setState(newState); 
  }
  //location='user.json'
  render() {
    return (
      <div>
        <Router>
          <div>
            <Nav user={this.state.user} handleSignedOut={this.handleSignedOut} showModalWindow={this.showSignInModalWindow}/>
            <div className='container pt-4 mt-4'>
              <Route exact path="/" render={() => <CardContainer location='/products' showBuyModal={this.showBuyModalWindow}/>} />
            </div>
          </div>
        </Router>
      </div>
    );
  }
}

export default App;
