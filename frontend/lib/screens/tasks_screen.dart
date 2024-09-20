import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:frontend/models/api_return.dart';
import 'package:frontend/screens/login_screen.dart';
import 'package:http/http.dart' as http;
import 'package:frontend/global.dart';
import 'package:frontend/models/task.dart';

class TasksScreen extends StatefulWidget {
  const TasksScreen({Key? key}) : super(key: key);

  @override
  State<TasksScreen> createState() => _TasksScreenState();
}

class _TasksScreenState extends State<TasksScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("タスク一覧"),
        actions: [
          IconButton(icon: const Icon(Icons.filter_list), onPressed: () {}),
          IconButton(onPressed: () {}, icon: const Icon(Icons.sell))
        ],
      ),
      body: FutureBuilder(
        future: getTaskList(),
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator());
          } else if (snapshot.hasData) {
            var data = snapshot.data as ApiReturn;

            if (data.statusCode != 200) {
              if (data.statusCode == 401) {
                WidgetsBinding.instance.addPostFrameCallback(
                  (_) {
                    Navigator.pushReplacement(
                      context,
                      MaterialPageRoute(
                          builder: (context) => const LoginScreen()),
                    );
                  },
                );
                return const Center(child: Text("Unauthorized"));
              } else {
                return const Center(child: Text("Error"));
              }
            } else {
              return ListView.builder(
                itemCount: data.body.length,
                itemBuilder: (context, index) {
                  return GestureDetector(
                    child: Card(
                      child: ListTile(
                        title: Text(data.body[index].title),
                        leading: Checkbox(value: false, onChanged: (value) {}),
                        trailing: Text(data.body[index].deadline),
                      ),
                    ),
                    onTap: () {
                      showModalBottomSheet(
                        context: context,
                        builder: (context) {
                          return Container(
                            height: 200,
                            width: MediaQuery.sizeOf(context).width,
                            padding: const EdgeInsets.all(16),
                            child: Column(
                              mainAxisAlignment: MainAxisAlignment.center,
                              mainAxisSize: MainAxisSize.min,
                              children: [
                                Text(data.body[index].title),
                                Text(data.body[index].memo),
                                Text(data.body[index].deadline),
                              ],
                            ),
                          );
                        },
                      );
                    },
                  );
                },
              );
            }
          } else {
            return const Center(child: Text("Error"));
          }
        },
      ),
    );
  }

  Future<ApiReturn> getTaskList() async {
    if (Token == "") {
      return ApiReturn(statusCode: 401, body: "Token is empty");
    }

    if (UserID == "") {
      await getUserId().then((value) {
        if (value.statusCode == 200) {
          UserID = value.body;
        } else if (value.statusCode == 401) {
          return ApiReturn(statusCode: 401, body: "Token is empty");
        } else {
          return ApiReturn(statusCode: value.statusCode, body: "Error");
        }
      });
    }

    final query = {
      'user_id': UserID,
    };
    final header = <String, String>{
      'Authorization': 'Bearer $Token',
    };
    final url = Uri.parse(
        '$BaseURL/api/v1/auth/tasks?${Uri(queryParameters: query).query}');
    print(url);
    final response = await http.get(url, headers: header);

    print(response.statusCode);

    if (response.statusCode == 200) {
      final List<dynamic> body = jsonDecode(response.body);
      print(body);
      return ApiReturn(
          statusCode: 200,
          body: body.map((dynamic json) => Task.fromJson(json)).toList());
    } else if (response.statusCode == 401) {
      return ApiReturn(statusCode: 401, body: "Unauthorized");
    } else {
      return ApiReturn(statusCode: response.statusCode, body: "Error");
    }
  }
}
