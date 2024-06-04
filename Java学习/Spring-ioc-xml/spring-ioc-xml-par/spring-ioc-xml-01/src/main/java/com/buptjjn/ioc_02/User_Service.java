package com.buptjjn.ioc_02;

public class User_Service {
    private User_DAO userDao;
    private int age;
    private String name;

    public User_Service() {

    }


    public User_Service(User_DAO userDao) {
        this.userDao = userDao;
    }

    public User_Service(User_DAO userDao, int age, String name) {
        this.userDao = userDao;
        this.age = age;
        this.name = name;
    }

    public void setAge(int age) {
        this.age = age;
    }

    public void setName(String name) {
        this.name = name;
    }
    public void setUserDao(User_DAO userDao) {
        this.userDao = userDao;
    }
}
