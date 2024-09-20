import 'dart:convert';

import 'package:frontend/models/api_return.dart';
import 'package:http/http.dart' as http;
import 'package:flutter/material.dart';

// トークンを格納するグローバル変数
String Token = "";
// ユーザーIDを格納するグローバル変数
String UserID = "";
// ユーザー名を格納するグローバル変数
String BaseURL = "http://192.168.1.22:8080";

Future<ApiReturn> getUserId() async {
  final header = <String, String>{
    'Content-Type': 'application/json',
    'Authorization': 'Bearer $Token',
  };
  print(header);
  final url = Uri.parse('$BaseURL/api/v1/auth/current_user');
  final response = await http.get(url, headers: header);

  print(url);
  print(response.statusCode);
  if (response.statusCode != 200) {
    return ApiReturn(statusCode: response.statusCode, body: "Error");
  }

  final user_id = jsonDecode(response.body)["user_id"];

  return ApiReturn(statusCode: response.statusCode, body: user_id);
}
