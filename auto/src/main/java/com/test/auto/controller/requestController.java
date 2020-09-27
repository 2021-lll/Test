package com.test.auto.controller;

import com.test.auto.model.HttpRequestForm;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
//import lombok.extern.log4j.Log4j;
import org.mybatis.spring.SqlSessionTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;

//@Log4j
@RestController
@Api(value = "v1")
@RequestMapping("v1")
public class requestController {
    @Autowired
    private SqlSessionTemplate template;

    @ApiOperation(value = "添加请求接口", httpMethod = "POST")
    @RequestMapping(value = "/addapi", method = RequestMethod.POST)
    public boolean addApi( @RequestBody HttpRequestForm httpRequestForm){
        int result = 0;
        result = template.insert("addR",httpRequestForm);
        if (result==1){
            return true;
        }
        return false;
    }

}
