package com.test.auto.testcase;

import com.test.auto.config.TestConfig;
import com.test.auto.model.HttpRequestForm;
import com.test.auto.util.DataBaseUtil;
import org.apache.http.HttpEntity;
import org.apache.http.NameValuePair;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.entity.StringEntity;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.message.BasicNameValuePair;
import org.apache.http.util.EntityUtils;
import org.apache.ibatis.session.SqlSession;
import org.testng.annotations.Test;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;

import java.io.IOException;
import java.io.UnsupportedEncodingException;
import java.util.ArrayList;
import java.util.List;

public class addApiRequest {

    @Test( description = "添加api接口")
    public void test1() throws IOException, InterruptedException {
        SqlSession session = DataBaseUtil.getSqlSession();
        HttpRequestForm httpRequestForm = session.selectOne("httpRequestForm",1);
        System.out.println(httpRequestForm.toString());
        System.out.println(TestConfig.httpRequestUrl);
        String result = getResult(httpRequestForm);

        Thread.sleep(2000);
//        HttpRequestForm httpRequestForm1 = session.selectOne("httpRequestForm1",httpRequestForm);
//        System.out.println();
//        return null;
    }

    private String getResult(HttpRequestForm httpRequestForm) {
        String result =  "";
        CloseableHttpClient client = HttpClients.createDefault();
        HttpPost post = new HttpPost(TestConfig.httpRequestUrl);
        JSONObject param = new JSONObject();
        param.put("groupName",httpRequestForm.getGroupName());
        param.put("apiName",httpRequestForm.getApiName());
        param.put("urlAddress",httpRequestForm.getUrlAddress());
//        List<NameValuePair> parameters = new ArrayList<NameValuePair>(0);
//        parameters.add(new BasicNameValuePair("groupName",httpRequestForm.getGroupName()));
//        parameters.add(new BasicNameValuePair("apiName", httpRequestForm.getApiName()));
//        parameters.add(new BasicNameValuePair("urlAddress", httpRequestForm.getUrlAddress()));
        post.setHeader("Content-Type","application/json");
        try {
            post.setEntity(new StringEntity(param.toString(),"utf-8"));
            CloseableHttpResponse response = client.execute(post);
            try {
                HttpEntity entity = response.getEntity();
                result = EntityUtils.toString(entity,"utf-8");
                System.out.println(result);
            }
            finally {
                response.close();
            }
        }
        catch (ClassCastException e){
            e.printStackTrace();
        }
        catch (UnsupportedEncodingException e){
            e.printStackTrace();
        }
        catch (IOException e){
            e.printStackTrace();
        }finally {
            try {
                client.close();
            }
            catch (Exception e){
                e.printStackTrace();
            }
        }
        return result;
    }
}
