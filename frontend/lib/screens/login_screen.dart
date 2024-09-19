import 'package:flutter/material.dart';
import 'dart:convert';
import 'package:flutter/services.dart';
import 'package:http/http.dart' as http;
import 'package:frontend/global.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({super.key});

  @override
  State<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  bool _isVisible = false;
  String _userID = "";
  String _password = "";

  final _formKey = GlobalKey<FormState>();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Login'),
      ),
      body: SafeArea(
        child: Center(
          child: Container(
            padding: const EdgeInsets.all(16),
            child: Form(
              key: _formKey,
              child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: <Widget>[
                    TextFormField(
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return 'IDを入力してください。';
                        }
                        return null;
                      },
                      inputFormatters: [
                        FilteringTextInputFormatter.allow(
                            RegExp(r'[a-zA-Z0-9]')),
                        LengthLimitingTextInputFormatter(10),
                      ],
                      decoration: const InputDecoration(
                        labelText: 'ID',
                      ),
                      onSaved: (value) {
                        _userID = value!;
                      },
                    ),
                    TextFormField(
                        validator: (value) {
                          if (value == null || value.isEmpty) {
                            return 'パスワードを入力してください。';
                          }
                          if (value.length < 8) {
                            return 'パスワードは8文字以上で入力してください。';
                          }
                          return null;
                        },
                        obscureText: !_isVisible,
                        inputFormatters: [
                          FilteringTextInputFormatter.allow(
                              RegExp(r'[a-zA-Z0-9]')),
                          LengthLimitingTextInputFormatter(10),
                        ],
                        decoration: InputDecoration(
                            labelText: 'パスワード',
                            suffixIcon: IconButton(
                              onPressed: () {
                                setState(() {
                                  _isVisible = !_isVisible;
                                });
                              },
                              icon: Icon(_isVisible
                                  ? Icons.visibility
                                  : Icons.visibility_off),
                            )),
                        onSaved: (value) {
                          _password = value!;
                        }),
                    Center(
                      child: ElevatedButton(
                        onPressed: () {
                          if (_formKey.currentState!.validate()) {
                            _formKey.currentState!.save();
                            login(_userID, _password).then((value) {
                              final data = jsonDecode(value);
                              Token = data['token'];
                              ScaffoldMessenger.of(context).showSnackBar(
                                  const SnackBar(
                                      content: Text('ログインに成功しました。')));
                            });
                          }
                        },
                        child: const Text('ログイン'),
                      ),
                    ),
                  ]),
            ),
          ),
        ),
      ),
    );
  }

  Future<String> login(String id, String password) async {
    final body = jsonEncode(<String, String>{
      'name': id,
      'password': password,
    });
    final header = <String, String>{
      'Content-Type': 'application/json',
    };
    final url = Uri.parse('$BaseURL/login');
    final response = await http.post(url, body: body, headers: header);

    print(url);
    print(response.statusCode);

    if (response.statusCode == 200) {
      return response.body;
    } else {
      throw Exception('Failed to login');
    }
  }
}
