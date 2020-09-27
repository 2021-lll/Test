package com.test.auto;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;;
import org.springframework.scheduling.annotation.EnableScheduling;


@EnableScheduling
@SpringBootApplication
//@Controller
//@ComponentScan
public class AutoApplication {

//    @RequestMapping("/")
//    @ResponseBody
//    String home(){
//        return "hellow word";
//    }

    public static void main(String[] args) {
        SpringApplication.run(AutoApplication.class, args);
    }

}
