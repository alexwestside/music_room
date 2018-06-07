//
//  RegisterPageViewController.swift
//  MusicRoom
//
//  Created by Maria Aksenova on 6/7/18.
//  Copyright © 2018 Maria Aksenova. All rights reserved.
//

import UIKit

class RegisterPageViewController: UIViewController {

    @IBOutlet weak var userEmailTextField: UITextField!
    @IBOutlet weak var userPasswordTextField: UITextField!
    @IBOutlet weak var repeatPasswordTextField: UITextField!
    
    override func viewDidLoad() {
        super.viewDidLoad()

        // Do any additional setup after loading the view.
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }
    
    
//    @IBAction func registerButtonTapped(_ sender: Any) {
//        
//        let userEmail: String? = userEmailTextField.text;
//        let userPassword: String? = userPasswordTextField.text;
//        let userRepeatPassword: String? = repeatPasswordTextField.text;
//        
//        // Check for empty fields
//        if (userEmail!.isEmpty || userPassword!.isEmpty || userRepeatPassword!.isEmpty)
//        {
//            // Display alert message
//            displayMyAlertMessage(userMessage: "All fields are required");
//            return;
//            
//        }
//        
//        // Check if passwords match
//        if (userRepeatPassword != userPassword)
//        {
//            // Display alert message
//            displayMyAlertMessage(userMessage: "Password do not match");
//            return;
//        }
//    
//        
//        // Stored data
//        
//        //Data alert message with conformation
//    }
//    
//    func displayMyAlertMessage(userMessage:String)
//    {
//        let myAlert = UIAlertController(title:"Alert", message: userMessage, preferredStyle: UIAlertControllerStyle.alert);
//        
//        let okAction = UIAlertAction(title:"Ok", style: UIAlertActionStyle.default, handler: nil);
//        
//        myAlert.addAction(okAction);
//        
//        self.present(_: myAlert, animated: true, completion: nil);
//        
//    }
    
    /*
    // MARK: - Navigation

    // In a storyboard-based application, you will often want to do a little preparation before navigation
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        // Get the new view controller using segue.destinationViewController.
        // Pass the selected object to the new view controller.
    }
    */

}
