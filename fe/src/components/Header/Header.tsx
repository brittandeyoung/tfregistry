import React, { FC, useState } from 'react';
import logo from '../../static/logo.png';
import { HeaderWrapper } from './Header.styled';

interface HeaderProps { }

const Header: FC<HeaderProps> = () => (
   <HeaderWrapper data-testid="Header">
      <header>
         <div className="ui fixed borderless huge menu">
            <a className="header item" href="/">
               <img src={logo} className="tfregistry-logo" alt="logo" />
            </a>
            <div className="ui container grid">
               <div className="computer only row">
                  <div className="right menu">
                     <div className="header item">
                        <a href="https://docs.tfmodule.com"
                           data-tooltip="Documentation Site"
                           data-position="bottom center"><i className="book icon"></i></a>
                     </div>
                     <div className="header item">
                        <a href="mailto: feedback@deyoung.dev"
                           data-tooltip="Submit Feedback"
                           data-position="bottom center"><i className="question circle outline icon"></i></a>
                     </div>
                     <div className="header item">
                        <a className="ui button" href="{% url 'account_signup' %}">Register</a>
                     </div>
                     <div className="header item">
                        <a className="ui button" href="{% url 'account_login' %}">Log in</a>
                     </div>
                  </div>
               </div>
            </div>
         </div>
      </header>
   </HeaderWrapper>
);

export default Header;
