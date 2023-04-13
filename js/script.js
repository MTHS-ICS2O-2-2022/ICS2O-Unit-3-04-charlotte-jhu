// Copyright (c) 2023 Charlotte Jhu All rights reserved
//
// Created by: Charlotte Jhu
// Created on: April 2023
// This file contains the JS functions for index.html

'use strict'

/**
 * This function calculates degrees Fahrenheit to degrees Celsius
 */

function myButtonClicked() {
  // input
  const fahrenheit = parseFloat(document.getElementById('fahrenheit').value)

  // process
  const celsius = (fahrenheit - 32) * 5 / 9

  // output
  document.getElementById('celsius').innerHTML = "Tempurature is " + celsius + "°C"
}
