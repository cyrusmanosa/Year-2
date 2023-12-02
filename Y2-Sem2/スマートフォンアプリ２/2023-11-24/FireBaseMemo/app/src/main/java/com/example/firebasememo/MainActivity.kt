package com.example.firebasememo

import android.os.Bundle
import android.util.Log
import androidx.appcompat.app.AppCompatActivity
import androidx.appcompat.widget.Toolbar
import com.example.firebasememo.databinding.ActivityMainBinding
import com.google.firebase.firestore.ktx.firestore
import com.google.firebase.ktx.Firebase

class MainActivity : AppCompatActivity() {

    // ActivityMainBindingは、このアクティビティで使用するレイアウトのビューバインディングです。
    // ビューバインディングは、レイアウトXMLのビューに直接アクセスできるようにする機能です。
    private lateinit var binding: ActivityMainBinding

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        binding = ActivityMainBinding.inflate(layoutInflater)
        setContentView(binding.root)

        // アプリのツールバーを設定します。
        val toolbar: Toolbar = binding.appToolbar
        setSupportActionBar(toolbar)

    }
}

