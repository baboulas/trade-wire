// sign up and log in page

import React from 'react';

class EntryForm extends Component {
  render(){
    return (
      <div>
        {SignUpViewContent view = signup}
      </div>
    )
  }
}

const SignUpViewContent = {
  switch (view) {
    case logIn:
      return 
      （
        //log in form component
        <form onSubmit={this.onSubmitHandler}>
          <input type="email">email</input>
          <input type="password">password></input>
        </form>; 
        )
    case signUp:
      return <div className="sign-up">Sign up form</div>; // sign up form component
    default:
      return {<div className="log-in">Log in form</div>;} //default to log in form
  }
}

const SwitcherComponent extends Component {
  render(){
    return (
      <div className="switcher">
        <span>Log In</span>
        <span>Sign Up</span>
      </div>
    )
  }
}

export default EntryForm;
